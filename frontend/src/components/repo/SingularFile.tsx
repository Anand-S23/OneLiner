"use client";

import Editor, { Monaco } from '@monaco-editor/react';
import * as monaco from 'monaco-editor';
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ChangeEvent, MutableRefObject, useRef, useState } from 'react';
import { Trash2 } from 'lucide-react';

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
    readOnly: false,
    cursorStyle: 'line',
    automaticLayout: true,
};

interface SingularFileProps {
    index: number;
    deleteable: boolean;
    setFilename: (index: number, filename: string) => void;
    setEditorRef: (index: number, editorRef: MutableRefObject<monaco.editor.IStandaloneCodeEditor | null>) => void;
    deleteFile: (index: number) => void;
}

type LanguageMapType = {
    [ext: string]: string
}

// TODO: Add more languages
const LanguageMap: LanguageMapType = {
    "js": "javascript",
    "ts": "typescript",
    "py": "python",
    "c": "c",
    "cpp": "cpp",
    "bat": "bat",
    "css": "css",
    "scss": "scss",
    "json": "json",
    "html": "html",
    "xml": "xml",
    "php": "php",
    "cs": "csharp",
    "md": "markdown",
    "go": "go",
    "java": "java",
    "lua": "lua",
}

const SingularFile = (props: SingularFileProps) => {
    const editorRef = useRef<monaco.editor.IStandaloneCodeEditor | null>(null);
    const [ext, setExt] = useState<string>('');

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
        let filename = e.target.value;

        var nameArr = filename.trim().split(".");
        let extenstion = "";
        if (nameArr[0] !== "" && nameArr.length === 2)  {
            extenstion = nameArr.pop() ?? "";
        }

        if (ext !== extenstion) {
            const detectedLanguage = LanguageMap[extenstion] ?? "none";
            changeEditorLanguage(detectedLanguage);
            setExt(extenstion);
            console.log(detectedLanguage);
        }

        props.setFilename(props.index, e.target.value);
    }

    return (
        <div className='p-4'>
            <div className='border-black'>
                <div className='flex justify-between'>
                    <Input 
                        type="text" id="filename" name="filename" placeholder="Filename with extenstion"
                        className="p-4 w-full focus-visible:ring-offset-0"
                        onChange={(e) => handleFilenameUpdate(e)}
                    />

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
            </div>

            <div className='h-80'>
                <Editor 
                    height="100%" 
                    defaultLanguage="none"
                    defaultValue="" 
                    theme="vs-dark"
                    options={options}
                    onMount={handleEditorDidMount}
                />
            </div>
        </div>
    );
}

export default SingularFile;
