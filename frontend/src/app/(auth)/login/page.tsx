'use client';

import LoginForm from "@/components/auth/LoginForm";
import StaticNavbar from "@/components/ui/static-navbar";
import { useRouter } from "next/navigation";

function Login() {
    const router = useRouter();

    return (
        <>
            <StaticNavbar
                text="Register"
                redirect="/register"
            />

            <div className="flex justify-center items-center h-full py-12">
                <div className="bg-white p-8 rounded shadow-lg w-96">
                    <h1 className="px-5 text-2xl text-semibold pt-5"> Login </h1>

                    <LoginForm />

                    <p className="px-5">
                        New User? <span 
                            className="text-blue-500 hover:text-blue-400 hover:cursor-pointer hover:underline"
                            onClick={() => router.push('/register')}
                        >Register</span>
                    </p> 
                </div>
            </div>
        </>
    );
}

export default Login;
