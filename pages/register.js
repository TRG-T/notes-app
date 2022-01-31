import styles from "../styles/Login.module.css";
import axios from "axios"
import Link from 'next/link'
import { useFormik } from "formik";
import Router from 'next/router';
import { useContext, useState } from "react";
import { AuthContext } from "./_app";

export default function Register() {
  const Auth = useContext(AuthContext)
  const [response, setResponse] = useState();
  const formik = useFormik({
    initialValues: {
      username: "",
      password: "",
    },
    onSubmit: ({ username, password }) => {
      axios.post("http://localhost:3001/register", {username, password}).then(response => { 
        setResponse(response.data)
        if (response.data === "Success") {
          Router.push("/home")
          Auth.setAuth(true)
          Auth.setUser(username)
        }
      })
    }
  });

  return (
    <div className={styles.container}>
      <div className={styles.login}>
        <h2>Register</h2>
        <form onSubmit={formik.handleSubmit} className={styles.form}>
          <input id="username" name="username" placeholder="username" type="text" onChange={formik.handleChange} value={formik.values.username} />
          <input id="password" name="password" placeholder="password" type="password" onChange={formik.handleChange} value={formik.values.password} />
          <p className={styles.response}>{response}</p>
          <button type="submit" className={styles.button}>Register</button>
        </form>
        <p className={styles.register}>
          Have an account? <Link href="/">Log in</Link>
        </p>
      </div>
    </div>
  );
}
