import { Flex } from "@chakra-ui/react";
import { ApiButtonUser } from "../components/apiButtonUser";


export default function Home() {

    return (
        <div>
            <Flex direction={"column"} justify={"center"} align="center" height="100vh">
                <ApiButtonUser />
            </Flex>
            <footer>

            </footer>
        </div>
    );
}
