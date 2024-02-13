'use client'

import { useRouter } from "next/navigation";
import { AUTH_USER_ENDPOINT, LOGOUT_ENDPOINT } from "@/lib/consts";
import { useEffect, useState } from "react";

const Navbar = () => {
    const router = useRouter();

    const [userID, setUserID] = useState<string>('');
    const [isLoaded, setIsLoaded] = useState<boolean>(false);

    useEffect(() => {
        const doAuth = async () => {
            const response = await fetch(AUTH_USER_ENDPOINT, {
                method: "GET",
                mode: "cors",
                headers: { "Content-Type": "application/json" },
                credentials: 'include'
            });

            if (!response.ok) {
                setUserID('');
            } else {
                const userID: string = await response.json();
                setUserID(userID);
            }

            setIsLoaded(true);
        }

        doAuth();
    }, []);


    const logout = async () => {
        await fetch(LOGOUT_ENDPOINT, {
            method: "POST",
            mode: "cors",
            headers: { "Content-Type": "application/json" },
            credentials: 'include'
        });

        router.push('/login');
    }

    if (!isLoaded) {
        return (<></>);
    }

    return (
        <nav className="bg-black px-8 md:px-16 lg:px-40 p-4 text-white flex justify-between items-center">
            <div 
                className="text-xl font-bold hover:cursor-pointer"
                onClick={() => router.push('/')}
            >
                {'<Snippet />'}
            </div>

            { userID === '' &&
                <p 
                    className="text-white text-lg hover:cursor-pointer hover:underline px-2"
                    onClick={() => router.push('/login')}
                >
                    {window.location.pathname !== '/login' ? 'Login' : 'Register'}
                </p>
            }

            { userID !== '' && 
                <p 
                    className="text-white text-lg hover:cursor-pointer hover:underline"
                    onClick={logout}
                >
                    Logout
                </p>
            }
        </nav>
    );
}

export default Navbar;
