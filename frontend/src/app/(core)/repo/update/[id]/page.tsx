'use client';

import UpdateForm from "@/components/repo/UpdateForm";

function Update({ params }: { params: { id: string } }) {
    return (
        <UpdateForm 
            repoID={params.id}
        />
    );
}

export default Update;
export const runtime = 'edge';
