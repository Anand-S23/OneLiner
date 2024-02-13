'use client';

import RegisterForm from "@/components/auth/RegisterForm";
import StaticNavbar from "@/components/ui/static-navbar";

function Register() {
    return (
        <>
            <StaticNavbar
                text="Login"
                redirect="/login"
            />

            <RegisterForm />
        </>
    );
}

export default Register;
