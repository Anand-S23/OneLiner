'use client';

import RegisterForm from "@/components/auth/RegisterForm";
import StaticNavbar from "@/components/ui/static-navbar";
import { useRouter } from "next/navigation";

function Register() {
    const router = useRouter();

    return (
        <>
            <StaticNavbar
                text="Login"
                redirect="/login"
            />

            <div className="flex justify-center items-center h-full py-12">
                <div className="bg-white p-8 rounded shadow-lg w-96">
                    <h1 className="px-5 text-2xl text-semibold pt-5"> Register </h1>

                    <RegisterForm />

                    <p className="px-5">
                        Already a User? <span 
                            className="text-blue-500 hover:text-blue-400 hover:cursor-pointer hover:underline"
                            onClick={() => router.push('/login')}
                        >Login</span>
                    </p> 
                </div>
            </div>
        </>
    );
}

export default Register;
