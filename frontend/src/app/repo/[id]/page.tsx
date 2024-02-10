'use client';

import { GET_FILES_ENDPOINT, READ_REPO_ENDPOINT } from "@/lib/consts";
import { FileDetails, FilesType, Post } from "@/lib/types";
import { useRouter } from "next/navigation";
import { MutableRefObject, useEffect, useState } from "react";
import * as monaco from 'monaco-editor';
import SingularFile from "@/components/repo/SingularFile";

function Read({ params }: { params: { id: string } }) {
    const router = useRouter();
    const [post, setPost] = useState<Post | null> (null);
    const [isLoaded, setIsLoaded] = useState<boolean>(false);

    const [fileContent, setFileContent] = useState<FilesType>({});
    const [files, setFiles] = useState<Array<FileDetails>>([]);

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
            setPost(data);

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

            for (let key in filesContent) {
                console.log(key);
                setFiles([...files, {
                    name: key, editorRef: null
                }]);
            }

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
        <div>
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
        </div>
    );
}

export default Read;
