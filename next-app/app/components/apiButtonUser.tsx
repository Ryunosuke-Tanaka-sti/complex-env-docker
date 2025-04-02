"use client"

import { Button } from "@chakra-ui/react";
import { useState } from "react";

export const ApiButtonUser = () => {
    const [res, setRes] = useState<string>()
    const onButtonClick = async () => {
        console.log("Button clicked");
        const resposen = await fetch("/api/user")
        const data = await resposen.json()
        setRes(data)
        console.log(data)
    };
    return (
        <>
            <Button px={8} py={2} onClick={onButtonClick}>API User</Button>
            <pre>
                {JSON.stringify(res, null, 2)}
            </pre>
        </>

    )

}