import React from "react";
import { Button } from "./button";

interface ModalProps {
    closeModal: () => void; 
    onConfirm: () => void; 
}

const Modal = ({ closeModal, onConfirm }: ModalProps) => {
    return (
        <>
        <div className="justify-center items-center flex overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none">
            <div className="relative w-auto my-6 mx-auto max-w-3xl">
                    <div className="border-0 rounded-lg shadow-lg relative flex flex-col w-full bg-white outline-none focus:outline-none">
                        <div className="relative p-6 flex-auto">
                            <p className="my-4 text-blueGray-500 text-lg leading-relaxed">
                                Are you sure you want to proceed?
                            </p>
                        </div>

                        <div className="flex items-center justify-center p-6 border-t border-solid border-blueGray-200 rounded-b">
                            <Button onClick={closeModal} variant={"outline"} className="mx-2">
                                Cancel
                            </Button>

                            <Button onClick={onConfirm} className="mx-2">
                                Confirm
                            </Button>
                        </div>
                    </div>
                </div>
            </div>

            <div className="opacity-25 fixed inset-0 z-40 bg-black"></div>
        </>
    ); 
}

export default Modal;
