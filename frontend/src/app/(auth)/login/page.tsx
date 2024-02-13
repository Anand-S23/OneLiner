'use client';

import LoginForm from "@/components/auth/LoginForm";
import StaticNavbar from "@/components/ui/static-navbar";

function Login() {
    return (
        <>
            <StaticNavbar
                text="Register"
                redirect="/register"
            />
            <LoginForm />
        </>
    );
}

export default Login;
