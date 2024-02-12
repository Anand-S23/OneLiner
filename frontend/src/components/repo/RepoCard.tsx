'use client';

import { DropdownMenu } from "@radix-ui/react-dropdown-menu";
import { Code, MoreVertical, Pencil, Trash2 } from "lucide-react";
import { useRouter } from "next/navigation";
import { DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "../ui/dropdown-menu";

interface RepoCardProps {
    name: string;
    description: string;
    repoID: string;
    deleteRepo: (repoID: string) => void;
}

const RepoCard = (props: RepoCardProps) => {
    const router = useRouter();

    return (
        <div className="border border-black rounded-sm">
            <div className="p-4">
                <div className='flex justify-between'>
                    <div
                        className='hover:cursor-pointer flex'
                        onClick={() => router.push(`/repo/view/${props.repoID}`)}
                    >
                        <Code className="mt-1"/>
                        <p className="text-blue-500 text-2xl px-2">{props.name}</p>
                    </div>

                    <DropdownMenu>
                        <DropdownMenuTrigger>
                            <MoreVertical />
                        </DropdownMenuTrigger>
                        <DropdownMenuContent>
                            <DropdownMenuItem
                                onClick={() => router.push(`/repo/update/${props.repoID}`)}
                            >
                                <Pencil />
                                Edit
                            </DropdownMenuItem>
                            <DropdownMenuItem
                                onClick={() => props.deleteRepo(props.repoID)}
                            >
                                <Trash2 />
                                Delete
                            </DropdownMenuItem>
                        </DropdownMenuContent>
                    </DropdownMenu>
                </div>

                <p className='py-2 text-lg'>{props.description}</p>
            </div>
        </div>
    );
}

export default RepoCard;
