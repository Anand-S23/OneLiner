'use client'

import { useRouter } from "next/navigation";

interface StaticNavbarProps {
    text: string;
    redirect: string;
}

const StaticNavbar = ({ text, redirect }: StaticNavbarProps) => {
    const router = useRouter();

    return (
        <nav className="bg-black px-8 md:px-16 lg:px-40 p-4 text-white flex justify-between items-center">
            <div 
                className="text-xl font-bold hover:cursor-pointer"
                onClick={() => router.push('/')}
            >
                {'<Snippet />'}
            </div>

            <p 
                className="text-white text-lg hover:cursor-pointer hover:underline px-2"
                onClick={() => router.push(redirect)}
            >
                { text }
            </p>
        </nav>
    );
}

export default StaticNavbar;
