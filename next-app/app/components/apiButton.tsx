"use client"

import { Button } from "@chakra-ui/react";

export const ApiButton = () => {
    const onButtonClick = async () => {
        console.log("Button clicked");
        const resposen = await fetch("http://localhost:8080/api/")
        const data = await resposen.json()
        console.log(data)
    };
    return <Button px={8} py={2} onClick={onButtonClick}>API Hello</Button>

}