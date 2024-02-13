'use client';

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { LoginSchema, TLoginSchema } from "@/lib/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import { cn } from "@/lib/utils";
import { LOGIN_ENDPOINT } from "@/lib/consts";

const LoginForm = () => {
    const router = useRouter();

    const {
        register,
        handleSubmit,
        formState: { errors },
        setError
    } = useForm<TLoginSchema>({
        resolver: zodResolver(LoginSchema),
    });

    const onSubmit = async (data: TLoginSchema) => {
         const response = await fetch(LOGIN_ENDPOINT, {
            method: "POST",
            mode: "cors",
            body: JSON.stringify(data),
            headers: { "Content-Type": "application/json" },
            credentials: 'include'
        });

        const resData = await response.json();

        if (!response.ok) {
            if (resData.error) {
                setError("root", {
                    type: "server", 
                    message: resData.error
                });
            }

            return;
        }

        // Redirect to home upon sucessful login
        router.push("/");
    }

    return (
        <div className="flex flex-col w-full max-w-md gap-1.5 p-5">
            <form onSubmit={handleSubmit(onSubmit)}>
                <Input 
                    {...register("email")}
                    type="email" id="email" name="email" placeholder="Email"
                    className={cn("mt-2", errors.email || errors.root ? 'border-red-500' : '')}
                />
                { errors.email && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.email.message}`}
                    </p>
                )}

                <Input 
                    {...register("password")}
                    type="password" id="password" name="password" placeholder="Password"
                    className={cn("mt-2", errors.password || errors.root ? 'border-red-500' : '')}
                />
                { errors.password && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.password.message}`}
                    </p>
                )}

                { errors.root && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.root.message}`}
                    </p>
                )}

                <Button type="submit" className="mt-2">Submit</Button>
            </form>
        </div>
    );
}

export default LoginForm;
