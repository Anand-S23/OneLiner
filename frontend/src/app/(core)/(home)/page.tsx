'use client';

import RepoCard from '@/components/repo/RepoCard';
import Modal from '@/components/ui/modal';
import { useToast } from '@/components/ui/use-toast';
import { DELETE_FILES_ENDPOINT, DELETE_REPO_ENDPOINT, POSTS_ENDPOINT } from '@/lib/consts';
import { FilesType, Post } from '@/lib/types';
import { Plus } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';

export default function Home() {
    const router = useRouter();
    const [posts, setPosts] = useState<Array<Post>>([]);
    const [isLoaded, setIsLoaded] = useState<boolean>(false);
    const [showModal, setShowModal] = useState<boolean>(false);
    const [deleteID, setDeleteID] = useState<string>('');
    const [files, setFiles] = useState<FilesType | null>(null);
    const [user, setUser] = useState<string>('');
    const [refresh, setRefresh] = useState<boolean>(false);

    const { toast } = useToast();

    useEffect(() => {
        const getPosts = async () => {
            const response = await fetch(POSTS_ENDPOINT, {
                method: "GET",
                mode: "cors",
                headers: { "Content-Type": "application/json" },
                credentials: 'include'
            });

            if (!response.ok) {
                toast({
                    title: "Uh oh! Something went wrong.",
                    description: "There was a problem with your request.",
                });

                setPosts([]);
                router.push("/login");
                return;
            }

            const data: Array<Post> = await response.json() as Array<Post> ?? [];
            setPosts(data as Array<Post>);
            setIsLoaded(true);
        }

        getPosts();
    }, [refresh]);

    const startDeleteRepo = (repoID: string, files: FilesType, userID: string) => {
        console.log(repoID);
        setDeleteID(repoID);
        setFiles(files);
        setUser(userID);
        setShowModal(true);
    }

    const deleteRepo = async () => {
        if (deleteID === '') {
            return;
        }

        const response = await fetch(DELETE_REPO_ENDPOINT + deleteID, {
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
        } else {
            const deleteData = {
                userID: user,
                files: files
            }

            await fetch(DELETE_FILES_ENDPOINT, {
                method: "POST",
                mode: "cors",
                body: JSON.stringify(deleteData),
                headers: { "Content-Type": "application/json" },
                credentials: 'include'
            });
        }

        setDeleteID('');
        setShowModal(false);
        setRefresh((prevRefresh) => !prevRefresh);
    }

    if (!isLoaded) {
        return <></>;
    }

    return (
        <>
            <div className="p-2 px-8 md:px-16 lg:px-40 grid grid-col-1 md:grid-col-2 lg:grid-col-3 gap-4">
                <div className='hover:cursor-pointer mt-2' onClick={() => router.push('/repo/create')}>
                    <div className='w-full flex justify-around align-middle border border-black rounded-sm hover:border-blue-500'>
                        <div className='flex my-5'>
                            <Plus className='mt-1'/>
                            <h2 className='text-2xl'>Create Repo</h2>
                        </div>
                    </div>
                </div>

                { posts.length === 0 &&
                    <p className='text-center text-lg'>
                        No posts created yet.
                    </p>
                }

                { posts.map(post => {
                    return (
                        <div key={post.id}>
                            <RepoCard 
                                name={post.name}
                                description={post.description}
                                repoID={post.id}
                                deleteRepo={() => startDeleteRepo(post.id, post.files, post.userID)}
                            />
                        </div>
                    );
                })}
            </div>

            { showModal && 
                <Modal 
                    closeModal={() => setShowModal(false)}
                    onConfirm={ async () => { await deleteRepo(); } }
                />
            }
        </>
    );
}
