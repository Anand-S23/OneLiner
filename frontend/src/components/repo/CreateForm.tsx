'use client';

import { FormEvent, MutableRefObject, useState } from 'react';
import SingularFile from './SingularFile';
import * as monaco from 'monaco-editor';
import { Button } from '../ui/button';
import { Input } from '../ui/input';
import { Plus } from 'lucide-react';


interface FileDetails {
    name: string;
    editorRef: MutableRefObject<monaco.editor.IStandaloneCodeEditor | null> | null;
}

const CreateForm = () => {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [files, setFiles] = useState<Array<FileDetails>>([{ name: "", editorRef: null }]);

    const addNewFile = (e: FormEvent) => {
        e.preventDefault();

        if (files.length == 5) {
            return;
        }

        const newFile: FileDetails = { 
            name: "", 
            editorRef: null
        };

        setFiles((prevFiles) => [...prevFiles, newFile]);
    }

    const updateFilename = (index: number, filename: string) => {
        // TODO: validate filename
        let updateFiles = [...files];
        updateFiles[index].name = filename;
        setFiles(updateFiles);
    }

    const updateEditorRef = (index: number, ref: MutableRefObject<monaco.editor.IStandaloneCodeEditor | null>) => {
        let updateFiles = [...files];
        updateFiles[index].editorRef = ref;
        setFiles(updateFiles);
    }

    const deleteFile = (index: number) => {
        const updatedFiles = files.filter((_, idx) => index !== idx);
        console.log(updatedFiles)
        setFiles(updatedFiles);
    }

    // TODO: Implement this function so it is hitting the endpoints
    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        const formData = new FormData();
        console.log(formData, name, description);
    };

    return (
        <div className='sm:px-5 md:px-12 lg:px-40'>
            <form>
                <div className='p-4'>
                    <Input 
                        type="text" id="name" name="name" placeholder="Repo Name"
                        className="p-4 my-4 w-full focus-visible:ring-offset-0"
                        onChange={(e) => setName(e.target.value)}
                    />

                    <Input 
                        type="text" id="description" name="description" placeholder="Description"
                        className="p-4 mt-4 w-full focus-visible:ring-offset-0"
                        onChange={(e) => setDescription(e.target.value)}
                    />
                </div>

                <div>
                    { files.map((_, index) => {
                        return (
                            <SingularFile
                                key={index}
                                index={index}
                                deleteable={files.length > 1}
                                setFilename={updateFilename}
                                setEditorRef={updateEditorRef}
                                deleteFile={deleteFile}
                            />
                                
                        );
                    })}

                    { files.length < 5 &&
                        <Button 
                            onClick={(e) => addNewFile(e)}
                            variant={"outline"}
                            className='mx-4 mb-4 mt-2 flex justify-evenly hover:pointer border-gray-400'
                        >
                            <Plus />
                            Add Page
                        </Button>
                    }
                </div>

                <Button 
                    onClick={(e) => handleSubmit(e)}
                    className='mx-4 mb-4 mt-2'
                >
                    Submit
                </Button>
            </form>
        </div>
    );
};

export default CreateForm;

