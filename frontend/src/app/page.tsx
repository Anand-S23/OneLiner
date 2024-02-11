'use client';

import RepoCard from '@/components/repo/RepoCard';
import { POSTS_ENDPOINT } from '@/lib/consts';
import { Post } from '@/lib/types';
import { Plus } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';

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
            <div className="p-2 px-8 md:px-16 lg:px-40 grid grid-col-1 md:grid-col-2 lg:grid-col-3 gap-4">

                <div className='hover:cursor-pointer' onClick={() => router.push('/repo/create')}>
                    <div className='w-full flex justify-around align-middle border border-black rounded-sm'>
                        <div className='flex my-5'>
                            <Plus className='mt-1'/>
                            <h2 className='text-2xl'>Create Repo</h2>
                        </div>
                    </div>
                </div>

                { posts.map(post => {
                    return (
                        <div key={post.id}>
                            <RepoCard 
                                name={post.name}
                                description={post.description}
                                repoID={post.id}
                            />
                        </div>
                    );
                })}
            </div>
        </div>
    );
}
