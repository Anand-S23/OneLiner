'use client';

import { DELETE_REPO_ENDPOINT, GET_FILES_ENDPOINT, READ_REPO_ENDPOINT } from "@/lib/consts";
import { FileDetails, FilesType, Post } from "@/lib/types";
import { useRouter } from "next/navigation";
import { MutableRefObject, useEffect, useState } from "react";
import * as monaco from 'monaco-editor';
import SingularFile from "@/components/repo/SingularFile";
import { DropdownMenu } from "@radix-ui/react-dropdown-menu";
import { DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import { MoreVertical, Pencil, Trash2 } from "lucide-react";
import { useToast } from "@/components/ui/use-toast";
import Modal from "@/components/ui/modal";

function Read({ params }: { params: { id: string } }) {
    const router = useRouter();
    const { toast } = useToast();
    const [isLoaded, setIsLoaded] = useState<boolean>(false);

    const [repo, setRepo] = useState<Post | null>(null);
    const [fileContent, setFileContent] = useState<FilesType>({});
    const [files, setFiles] = useState<Array<FileDetails>>([]);
    const [showModal, setShowModal] = useState<boolean>(false);

    const deleteRepo = async () => {
        const response = await fetch(DELETE_REPO_ENDPOINT + params.id, {
            method: "POST",
            mode: "cors",
            headers: { "Content-Type": "application/json" },
            credentials: 'include'
        });

        if (!response.ok) {
            toast({
                title: "Uh oh! Something went wrong.",
                description: "There was a problem with your request.",
            });
        }

        setShowModal(false);
        router.push('/');
    }

    useEffect(() => {
        const getPost = async () => {
            const response = await fetch(READ_REPO_ENDPOINT + params.id, {
                method: "GET",
                headers: { "Content-Type": "application/json" },
            });

            if (!response.ok) {
                router.push('/404');
                return;
            }

            const data: Post = await response.json() as Post;
            setRepo(data);

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
        }

        getPost();
    }, []);

    const updateEditorRef = (index: number, ref: MutableRefObject<monaco.editor.IStandaloneCodeEditor | null>) => {
        let updateFiles = [...files];
        updateFiles[index].editorRef = ref;
        setFiles(updateFiles);
    }

    if (!isLoaded) {
        return <></>;
    }

    return (
        <div className='sm:px-5 md:px-12 lg:px-40'>
            <div className='p-2 flex justify-between'>
                <div className='hover:cursor-pointer flex'>
                    <p className="text-2xl px-2">{repo?.name}</p>
                </div>

                <DropdownMenu>
                    <DropdownMenuTrigger>
                        <MoreVertical />
                    </DropdownMenuTrigger>
                    <DropdownMenuContent>
                        <DropdownMenuItem
                            onClick={() => router.push(`/repo/update/${repo?.id}`)}
                        >
                            <Pencil />
                            Edit
                        </DropdownMenuItem>
                        <DropdownMenuItem
                            onClick={() => setShowModal(true)}
                        >
                            <Trash2 />
                            Delete
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>

            { files.map((file, index) => {
                return (
                    <SingularFile
                        key={index}
                        filename={file.name}
                        editorValue={fileContent[file.name]}
                        index={index}
                        deleteable={false}
                        editable={false}
                        setFilename={() => {}}
                        setEditorRef={updateEditorRef}
                        deleteFile={() => {}}
                    />
                );
            })}

            { showModal && 
                <Modal 
                    closeModal={() => setShowModal(false)}
                    onConfirm={ async () => { await deleteRepo(); } }
                />
            }
        </div>
    );
}

export default Read;
