import { useContext, useState} from "react";
import styles from "../styles/Home.module.css"
import Router from "next/router";
import Link from "next/link"
import { AuthContext } from "./_app";

const Home = () => {
    const Auth = useContext(AuthContext)
    const [addNote, setAddNote] = useState()
    if (Auth.auth === false ) {
        Router.back()
    }
    return (
        <div className={styles.container}>
            <div className={styles.nav}>
                <p>{Auth.user}</p>
                <Link 
                    href="/" 
                    onClick={() => Auth.setAuth(false)}
                >
                    Logout
                </Link>
                <button 
                    onClick={() => setAddNote(!addNote)} 
                    style={{ transform: `rotate(${addNote ? "45deg": "0deg"})`}} 
                    className={styles.addNote}
                >
                    +
                </button>
            </div>
            <div className={styles.main}>
                { addNote && 
                    <form className={styles.addNoteForm}>
                        <input placeholder="title" />
                        <textarea placeholder="description" />
                        <button type="submit">Submit</button>
                    </form> 
                }
            </div>
        </div>
    )
}

export default Home;