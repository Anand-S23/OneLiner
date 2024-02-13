'use client';

import { useRouter } from "next/navigation";

function NotFound() {
    const router = useRouter();

    return (
        <div className="flex items-center justify-center h-screen">
            <div className="text-center">
                <h1 className="text-3xl">
                    404 Page Not Found
                </h1>

                <p
                    onClick={() => router.push('/')}
                    className="text-lg text-blue-500 hover:underline hover:text-blue-400 hover:cursor-pointer"
                >
                    Return to Home
                </p>
            </div>
        </div>
    );
}

export default NotFound;
