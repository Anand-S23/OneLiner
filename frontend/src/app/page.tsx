'use client';

import { POSTS_ENDPOINT } from '@/lib/consts';
import { useEffect, useState } from 'react';

type FilesType = {
    [key: string]: string
}

interface Post {
    ID: string,
    Name: string
    Description: string,
    Files: FilesType,
    UserID: string,
    CreatedAt: Date
}

export default function Home() {
    const [isLoaded, setIsLoaded] = useState(false);
    const [posts, setPosts] = useState<Post[]>([]);

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
                setIsLoaded(true);
            }

            const data = await response.json() as Post[];
            console.log(data);
            setPosts(data);
            setIsLoaded(true);
        }

        getPosts();
    }, []);

    if (isLoaded) {
        return '';
    }

    return (
        <div>
            { posts.map(post => {
                return (
                    <div key={post.ID}>
                        <p> {post.ID} </p>
                        <p> {post.Name} </p>
                        <p> {post.Description} </p>
                        <p> {post.CreatedAt.toString()} </p>
                    </div>
                );
            })}
        </div>
    );
}
