'use client';

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { TRegisterSchema, RegisterSchema } from "@/lib/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import { cn } from "@/lib/utils";
import { REGISTER_ENDPOINT } from "@/lib/consts";

const RegisterForm = () => {
    const router = useRouter();

    const {
        register,
        handleSubmit,
        formState: { errors },
        setError
    } = useForm<TRegisterSchema>({
        resolver: zodResolver(RegisterSchema),
    });

    const onSubmit = async (data: TRegisterSchema) => {
         const response = await fetch(REGISTER_ENDPOINT, {
            method: "POST",
            mode: "cors",
            body: JSON.stringify(data),
            headers: { "Content-Type": "application/json" },
        });

        const resData = await response.json();

        if (resData.email || resData.password) {
            if (resData.email) {
                setError("email", {
                    type: "server", 
                    message: resData.email
                })
            }

            if (resData.password) {
                setError("password", {
                    type: "server", 
                    message: resData.password
                })
            }

            return;
        }

        if (!response.ok) {
            // TODO: Use toast here instead
            alert("Server error, please try again");
            return;
        }

        // TODO: potentially go striaght to login
        // Redirect to homepage
        router.push("/");
    };

    return (
        <div className="flex flex-col w-full max-w-md gap-1.5 p-5">
            <form onSubmit={handleSubmit(onSubmit)}>
                <Input 
                    {...register("email")}
                    type="email" id="email" name="email" placeholder="Email"
                    className={cn("mt-2", errors.email ? 'border-red-500' : '')}
                />
                { errors.email && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.email.message}`}
                    </p>
                )}

                <Input 
                    {...register("password")}
                    type="password" id="password" name="password" placeholder="Password"
                    className={cn("mt-2", errors.password ? 'border-red-500' : '')}
                />
                { errors.password && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.password.message}`}
                    </p>
                )}

                <Input 
                    {...register("confirm")}
                    type="password" id="confirm" name="confirm" placeholder="Confirm Password"
                    className={cn("mt-2", errors.confirm ? 'border-red-500' : '')}
                />
                { errors.confirm && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.confirm.message}`}
                    </p>
                )}

                <Button type="submit" className="mt-2">Submit</Button>
            </form>
        </div>
   );
} 

export default RegisterForm;
