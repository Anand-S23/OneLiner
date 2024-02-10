'use client';

import RepoCard from '@/components/repo/RepoCard';
import { POSTS_ENDPOINT } from '@/lib/consts';
import { FilesType } from '@/lib/types';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';

interface Post {
    id: string,
    name: string
    description: string,
    files: FilesType,
    userID: string,
    createdAt: Date
}

export default function Home() {
    const router = useRouter();
    const [posts, setPosts] = useState<Array<Post>>([]);
    const [isLoaded, setIsLoaded] = useState<boolean>(false);

    useEffect(() => {
        const getPosts = async () => {
            const response = await fetch(POSTS_ENDPOINT, {
                method: "GET",
                mode: "cors",
                headers: { "Content-Type": "application/json" },
                credentials: 'include'
            });

            if (!response.ok) {
                // TODO: Handle error
                console.log("Error occured while getting posts for signed in user");
                setPosts([]);
                router.push("/login");
                return;
            }

            const data: Array<Post> = await response.json() as Array<Post>;
            setPosts(data as Array<Post>);
            setIsLoaded(true);
        }

        getPosts();
    }, []);

    if (!isLoaded) {
        return <></>;
    }

    return (
        <div>
            <h1> POSTS </h1>
            <div className="p-2">
                { posts.map(post => {
                    return (
                        <div className="container mx-auto mt-8">
                            <div 
                                className='grid grid-col-1 md:grid-col-2 lg:grid-col-3'
                                key={post.id}
                            >
                                <RepoCard 
                                    name={post.name}
                                    description={post.description}
                                    repoID={post.id}
                                />
                            </div>
                        </div>
                    );
                })}
            </div>
        </div>
    );
}
