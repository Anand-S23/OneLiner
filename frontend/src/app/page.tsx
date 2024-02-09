'use client';

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
                        <div key={post.id}>
                            <p> Post ID: {post.id} </p>
                            <p> User ID: {post.userID} </p>
                            <p> Name: {post.name} </p>
                            <p> Description: {post.description} </p>
                            <p> Created At: {post.createdAt.toString()} </p>
                            <p> Files: </p>
                            <div>
                                { Object.keys(post.files).map((key) => (
                                    <div key={key}>
                                        <p className='px-4'>{`- ${key}: ${post.files[key]}`}</p>
                                    </div>
                                ))}
                            </div>
                        </div>
                    );
                })}
            </div>
        </div>
    );
}
