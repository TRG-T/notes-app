import { useContext} from "react";
import Router from "next/router";
import Link from "next/link"
import { AuthContext } from "./_app";

const Home = () => {
    const Auth = useContext(AuthContext)
    console.log(Auth.auth)
    if (Auth.auth === false ) {
        Router.back()
    }
    return (
        <Link href="/" onClick={() => Auth.setAuth(false)}>Logout</Link>
    )
}

export default Home;