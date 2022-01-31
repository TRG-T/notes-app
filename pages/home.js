import { useContext, useState} from "react";
import styles from "../styles/Home.module.css"
import Router from "next/router";
import Link from "next/link"
import { useFormik } from "formik";
import { AuthContext } from "./_app";
import axios from "axios";

const Home = () => {
    const Auth = useContext(AuthContext);
    const [addNote, setAddNote] = useState();
    if (Auth.auth === false ) {
        Router.back()
    }
    const addNoteForm = useFormik({
        initialValues: {
          title: "",
          description: "",
          user: Auth.user
        },
        onSubmit: ({ title, description, user }) => {
          axios.post("http://localhost:3001/add-note", {title, description, user}).then(response => { 
            console.log(response.data)
            addNoteForm.resetForm();
          })
        }
      });
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
                    <form onSubmit={addNoteForm.handleSubmit} className={styles.addNoteForm}>
                        <input id="title" name="title" placeholder="title" value={addNoteForm.values.title} onChange={addNoteForm.handleChange} />
                        <textarea id="description" name="description" placeholder="description"value={addNoteForm.values.description} onChange={addNoteForm.handleChange} />
                        <button type="submit">Submit</button>
                    </form> 
                }
            </div>
        </div>
    )
}

export default Home;