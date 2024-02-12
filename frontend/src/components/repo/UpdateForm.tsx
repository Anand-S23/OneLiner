'use client';

import { FormEvent, MutableRefObject, useEffect, useState } from 'react';
import SingularFile from './SingularFile';
import * as monaco from 'monaco-editor';
import { Button } from '../ui/button';
import { Input } from '../ui/input';
import { Plus } from 'lucide-react';
import { GET_FILES_ENDPOINT, READ_REPO_ENDPOINT, UPDATE_REPO_ENDPOINT, UPLOAD_FILES_ENDPOINT } from '@/lib/consts';
import { CreateRepoSchema, FileDetails, FilesType, Post, TCreateRepoSchema } from '@/lib/types';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { cn } from '@/lib/utils';
import { useToast } from '../ui/use-toast';
import { useRouter } from 'next/navigation';

interface UpdateFormProps {
    repoID: string;
}

const UpdateForm = (props: UpdateFormProps) => {
    const {
        register,
        handleSubmit,
        formState: { errors },
        setError,
        clearErrors,
        setValue
    } = useForm<TCreateRepoSchema>({
        resolver: zodResolver(CreateRepoSchema),
    });

    const { toast } = useToast();
    const router = useRouter();

    const [repo, setRepo] = useState<Post | null>(null);
    const [files, setFiles] = useState<Array<FileDetails>>([]);
    const [fileContent, setFileContent] = useState<FilesType>({});

    const [isLoaded, setIsLoaded] = useState<boolean>(false);
    const [refresh, doRefresh] = useState(0);

    useEffect(() => {
        const getPost = async () => {
            const response = await fetch(READ_REPO_ENDPOINT + props.repoID, {
                method: "GET",
                headers: { "Content-Type": "application/json" },
            });

            if (!response.ok) {
                router.push('/404');
                return;
            }

            const data: Post = await response.json() as Post;
            setValue('name', data.name);
            setValue('description', data.description);

            const getFilesResponse = await fetch(GET_FILES_ENDPOINT, {
                method: "POST",
                mode: 'cors',
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(data.files)
            });

            if (!getFilesResponse.ok) {
                // TODO: Error handling
                return;
            }

            const filesContent: FilesType = await getFilesResponse.json();
            setFileContent(filesContent);

            let updateFiles: Array<FileDetails> = [];
            for (let key in filesContent) {
                updateFiles.push({name: key, editorRef: null});
            }
            setFiles(updateFiles);
            setIsLoaded(true);
            doRefresh(prev => prev + 1);
        }

        getPost();
    }, []);

    const addNewFile = (e: FormEvent) => {
        e.preventDefault();

        if (files.length == 5) {
            return;
        }

        const newFile: FileDetails = { 
            name: '', editorRef: null
        };
        setFiles((prevFiles) => [...prevFiles, newFile]);
    }

    const updateFilename = (index: number, filename: string) => {
        let updateFiles = [...files];
        updateFiles[index].name = filename;
        setFiles(updateFiles);
        clearErrors('root');
    }

    const updateEditorRef = (index: number, ref: MutableRefObject<monaco.editor.IStandaloneCodeEditor | null>) => {
        console.log("Editor Ref Set");
        let updateFiles = [...files];
        updateFiles[index].editorRef = ref;
        setFiles(updateFiles);
    }

    const deleteFile = (index: number) => {
        const updatedFiles = files.filter((_, idx) => index !== idx);
        console.log(updatedFiles)
        setFiles(updatedFiles);
    }

    const uploadFiles = async () => {
        const formData = new FormData();
        const seen = new Set();

        for (let i = 0; i < files.length; ++i) {
            const editorData = files[i].editorRef?.current?.getValue() ?? "";
            const filename = files[i].name.trim();

            if (editorData === "" || filename === "") {
                console.log("filename", filename, "editorData", editorData);
                setError("root", {
                    type: "value", 
                    message: "All filenames must be set and no snippet should be empty"
                });
                return undefined;
            } else if (seen.has(filename)) {
                setError("root", {
                    type: "value", 
                    message: "Need to have unqiue filenames for snippets"
                });
                return undefined;
            }

            seen.add(filename);
            const fileData = new Blob([editorData], {type: "text/plain"});
            const file = new File([fileData], filename);
            formData.append("files", file);
        }

         const uploadResponse = await fetch(UPLOAD_FILES_ENDPOINT, {
            method: "POST",
            mode: "cors",
            body: formData,
            credentials: 'include'
        });

        return uploadResponse;
    }

    const onSubmit = async (data: TCreateRepoSchema) => {
        const uploadResponse = await uploadFiles();
        if (uploadResponse === undefined) {
            return;
        } else if (!uploadResponse.ok) {
            toast({
                title: "Uh oh! Something went wrong.",
                description: "There was a problem with your request.",
            });
            return;
        }
        const uploadFilesReponse = await uploadResponse.json() as FilesType;

        const updateData = {
            name: data.name,
            description: data.description,
            files: uploadFilesReponse,
        }

         const updateResponse = await fetch(UPDATE_REPO_ENDPOINT + props.repoID, {
            method: "POST",
            mode: "cors",
            body: JSON.stringify(updateData),
            headers: { "Content-Type": "application/json" },
            credentials: 'include'
        });

        const updateResponseData = await updateResponse.json();

        if (!updateResponse.ok) {
            toast({
                title: "Uh oh! Something went wrong.",
                description: "There was a problem with your request.",
            });

            return;
        }

        // TODO: Delete old files
        router.push(`/repo/view/${props.repoID}`);
    };

    if (!isLoaded) {
        return <></>;
    }

    return (
        <div className='sm:px-5 md:px-12 lg:px-40'>
            <form onSubmit={handleSubmit(onSubmit)}>
                <div className='p-4 pb-0'>
                    <Input 
                        {...register("name")}
                        type="text" id="name" name="name" placeholder="Repo Name"
                        className={cn("mt-2", errors.name ? 'border-red-500' : '')}
                    />
                    { errors.name && (
                        <p className="text-sm text-red-500 mx-2">
                            {`${errors.name.message}`}
                        </p>
                    )}

                    <Input 
                        {...register("description")}
                        type="text" id="description" name="description" placeholder="Description"
                        className={cn("mt-2", errors.description ? 'border-red-500' : '')}
                    />
                    { errors.description && (
                        <p className="text-sm text-red-500 mx-2">
                            {`${errors.description.message}`}
                        </p>
                    )}
                </div>

                <div>
                    <div className={errors.root ? 'border border-red-500 mt-2' : ''}>
                        { files.map((file, index) => {
                            return (
                                <SingularFile
                                    key={index}
                                    filename={file.name}
                                    editorValue={fileContent[file.name]}
                                    index={index}
                                    deleteable={files.length > 1}
                                    editable={true}
                                    setFilename={updateFilename}
                                    setEditorRef={updateEditorRef}
                                    deleteFile={deleteFile}
                                    refresh={refresh}
                                />
                            );
                        })}
                    </div>
                    { errors.root && (
                        <p className="text-sm text-red-500 mx-2">
                            {`${errors.root.message}`}
                        </p>
                    )}

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

                <Button type="submit" className="m-4 mt-2">Submit</Button>
            </form>
        </div>
    );
};

export default UpdateForm;

