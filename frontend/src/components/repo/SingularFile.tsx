import Editor, { Monaco } from '@monaco-editor/react';
import * as monaco from 'monaco-editor';
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ChangeEvent, MutableRefObject, useEffect, useRef, useState } from 'react';
import { Copy, Download, Trash2 } from 'lucide-react';
import { getFileExtension } from '@/lib/utils';
import { useToast } from '../ui/use-toast';
import { LanguageMap } from '@/lib/languages';

const getOptions = (editable: boolean) => {
    const options: monaco.editor.IStandaloneEditorConstructionOptions = {
        autoIndent: 'full',
        contextmenu: false,
        fontFamily: 'monospace',
        fontSize: 14,
        lineHeight: 24,
        hideCursorInOverviewRuler: true,
        matchBrackets: 'always',
        minimap: {
            enabled: false,
        },
        scrollbar: {
            horizontalSliderSize: 4,
            verticalSliderSize: 18,
        },
        selectOnLineNumbers: true,
        roundedSelection: false,
        readOnly: (!editable),
        domReadOnly: (!editable),
        cursorStyle: 'line',
        automaticLayout: true,
    };

    return options;
}

interface SingularFileProps {
    // Content
    filename: string;
    editorValue: string;

    // Functionality
    index: number;
    deleteable: boolean;
    editable: boolean;
    refresh?: number;
    setFilename: (index: number, filename: string) => void;
    setEditorRef: (index: number, editorRef: MutableRefObject<monaco.editor.IStandaloneCodeEditor | null>) => void;
    deleteFile: (index: number) => void;
}

const SingularFile = (props: SingularFileProps) => {
    const editorRef = useRef<monaco.editor.IStandaloneCodeEditor | null>(null);
    const [ext, setExt] = useState<string>('');
    
    const { toast } = useToast();

    useEffect(() => {
        props.setEditorRef(props.index, editorRef);
    }, [props.refresh])

    function handleEditorDidMount(editor: monaco.editor.IStandaloneCodeEditor, _monaco: Monaco) {
        editorRef.current = editor;                                                                                                                                                                                        
        props.setEditorRef(props.index, editorRef);
    }

    const changeEditorLanguage = (language: string) => {
        let model = editorRef.current?.getModel();

        if (model !== null && model !== undefined) {
            monaco.editor.setModelLanguage(model, language);
        }
    }

    const handleFilenameUpdate = (e: ChangeEvent<HTMLInputElement>) => {
        let filename = e.target.value.trim();
        let extenstion = getFileExtension(filename);

        if (ext !== extenstion) {
            const detectedLanguage = LanguageMap[extenstion] ?? "none";
            changeEditorLanguage(detectedLanguage);
            setExt(extenstion);
            console.log(detectedLanguage); // TODO: For debugging need to remove
        }

        props.setFilename(props.index, filename);
    }

    const copyContent = () => {
        let fileContent = editorRef.current?.getValue() ?? "";
        navigator.clipboard.writeText(fileContent);
        toast({ description: "File content copied to clipboard." });
    }

    const downloadFile = () => {
        let fileContent = editorRef.current?.getValue() ?? "";
        const blob = new Blob([fileContent], { type: 'application/octet-stream' });
        const downloadLink = document.createElement('a');
        downloadLink.href = URL.createObjectURL(blob);
        downloadLink.download = props.filename;
        document.body.appendChild(downloadLink);
        downloadLink.click();
        document.body.removeChild(downloadLink);
    }

    return (
        <div className='px-4 py-2'>
            <div>
                { props.editable &&
                    <div className='flex justify-between'>
                        { props.filename === '' ? (
                            <Input 
                                type="text" id="filename" name="filename" placeholder="Filename with extenstion"
                                className="p-4 w-full focus-visible:ring-offset-0"
                                onChange={(e) => handleFilenameUpdate(e)}
                            />) : (
                            <Input 
                                type="text" id="filename" name="filename" placeholder="Filename with extenstion"
                                value={props.filename}
                                className="p-4 w-full focus-visible:ring-offset-0"
                                onChange={(e) => handleFilenameUpdate(e)}
                            />
                        )}

                        { props.deleteable &&
                            <Button 
                                className="w-10 h-10 p-2"
                                variant={"destructive"}
                                onClick={() => props.deleteFile(props.index)}
                            >
                                <Trash2 />
                            </Button>
                        }
                    </div>
                }

                { !props.editable &&
                    <div className='flex justify-between'>
                        <h1 className='text-lg py-2'>{ props.filename }</h1>
                        <div>
                            <Button
                                variant={'outline'}
                                className='p-2 w-10 h-10'
                                onClick={copyContent}
                            >
                                <Copy />
                            </Button>
                            <Button
                                variant={'outline'}
                                className='p-2 w-10 h-10'
                                onClick={downloadFile}
                            >
                                <Download />
                            </Button>
                        </div>
                    </div>
                }
            </div>

            <div className='h-80'>
                <Editor 
                    height="100%" 
                    defaultLanguage={LanguageMap[getFileExtension(props.filename)] ?? "none"}
                    defaultValue={props.editorValue}
                    theme="vs-dark"
                    options={getOptions(props.editable)}
                    onMount={handleEditorDidMount}
                />
            </div>
        </div>
    );
}

export default SingularFile;
