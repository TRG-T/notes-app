import { useContext} from "react";
import Router from "next/router";
import { AuthContext } from "./_app";

const Home = () => {
    const Auth = useContext(AuthContext)
    console.log(Auth.auth)
    if (Auth.auth === false ) {
        Router.back()
    }
    return (
        <button onClick={() => { Auth.setAuth(false); Router.back() }}>Logout</button>
       
    )
}

export default Home;