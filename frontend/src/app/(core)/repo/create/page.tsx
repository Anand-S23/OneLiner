'use client';

import dynamic from "next/dynamic";

const CreateForm = dynamic(() => import("@/components/repo/CreateForm"), {ssr: false})

function Create() {
    return (
        <CreateForm />
    );
}

export default Create;
