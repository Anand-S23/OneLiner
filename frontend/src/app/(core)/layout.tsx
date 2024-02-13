import Navbar from "@/components/ui/navbar";

export default function AuthLayout({ children, }: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <>
            <Navbar />
            <div>{children}</div>
        </>
    );
}
