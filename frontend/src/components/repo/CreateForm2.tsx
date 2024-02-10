'use client';

import { FormEvent, MutableRefObject, useState } from 'react';
import SingularFile from './SingularFile';
import * as monaco from 'monaco-editor';
import { Button } from '../ui/button';
import { Input } from '../ui/input';
import { Plus, Upload } from 'lucide-react';
import { CREATE_REPO_ENDPOINT, UPLOAD_FILES_ENDPOINT } from '@/lib/consts';
import { FilesType, RepoDescriptionSchema, RepoNameSchema } from '@/lib/types';
import { ZodError, z } from 'zod';
import { error } from 'console';


interface FileDetails {
    name: string;
    editorRef: MutableRefObject<monaco.editor.IStandaloneCodeEditor | null> | null;
    error: string;
}

const CreateForm = () => {
    const [name, setName] = useState<string>('');
    const [nameError, setNameError] = useState<string>(''); 

    const [description, setDescription] = useState<string>('');
    const [descriptionError, setDescriptionError] = useState<string>('');

    const [files, setFiles] = useState<Array<FileDetails>>([
        { name: '', editorRef: null, error: ''}]
    );

    const addNewFile = (e: FormEvent) => {
        e.preventDefault();

        if (files.length == 5) {
            return;
        }

        const newFile: FileDetails = { 
            name: '', editorRef: null, error: ''
        };
        setFiles((prevFiles) => [...prevFiles, newFile]);
    }

    const updateFilename = (index: number, filename: string) => {
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

    const isFormValid = (name: string, description: string, uploadFilesCount: number) => {
        let valid = true;

        try {
            RepoNameSchema.parse(name);
        } catch (error) {
            if (error instanceof ZodError) {
                let errMsg = error.errors[0].message;
                console.log(errMsg);
                setNameError(errMsg);
            } else {
                // TODO: toast for unexpected error
            }

            valid = false;
        }

        try {
            RepoDescriptionSchema.parse(description);
        } catch (error) {
            if (error instanceof ZodError) {
                let errMsg = error.errors[0].message;
                console.log(errMsg);
                setDescriptionError(errMsg);
            } else {
                // TODO: toast for unexpected error
            }

            valid = false;
        }

        if (uploadFilesCount === 0) {
            // TODO: Error handling
            for (let i = 0; i < files.length; ++i) {
                files[i].error = "Snippet must have a filename and content"
            }

            valid = false;
        }
        
        return valid;
    }

    const hasDuplicateFilenames = (details: Array<FileDetails>) => {
        const seen = new Set();

        for (let i = 0; i < details.length; ++i) {
            let currentName = details[i].name;
            if (seen.has(currentName)) {
                return true;
            }

            seen.add(currentName);
        }

        return false;
    }

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();

        if (hasDuplicateFilenames(files)) {
            // TODO: Error handling
            console.log("Multiple files cannot have the same name");
            return;
        }

        const formData = new FormData();
        for (let i = 0; i < files.length; ++i) {
            const editorData = files[i].editorRef?.current?.getValue() ?? "";
            const filename = files[i].name.trim();
            if (editorData === "" || filename === "") {
                // TODO: error handling
                console.log("Need to have data in editor and name set");
                return;
            }
            const fileData = new Blob([editorData], {type: "text/plain"});
            const file = new File([fileData], filename);
            formData.append("files", file);
        }

        if (!isFormValid(name, description, formData.getAll("files").length)) {
            console.log("Form is not valid");
            return;
        }

         const uploadResponse = await fetch(UPLOAD_FILES_ENDPOINT, {
            method: "POST",
            mode: "cors",
            body: formData,
            credentials: 'include'
        });

        const uploadResponseData = await uploadResponse.json();
        const filesUploadResponse = uploadResponseData as FilesType;

        if (!uploadResponse.ok) {
            // TODO: Handle Error
            console.log("Error Uploading Files to S3", uploadResponse);
            return;
        }

        const createData = {
            name: name,
            description: description,
            files: filesUploadResponse
        }

         const createResponse = await fetch(CREATE_REPO_ENDPOINT, {
            method: "POST",
            mode: "cors",
            body: JSON.stringify(createData),
            headers: { "Content-Type": "application/json" },
            credentials: 'include'
        });

        const createResponseData = await createResponse.json();

        if (!createResponse.ok) {
            // TODO: Handle Error
            console.log("Could not create repo");
            return;
        }

        console.log(createResponseData);
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
                                filename=''
                                editorValue=''
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

