import styles from "../styles/Login.module.css";
import axios from "axios"
import { useFormik } from "formik";
import Router from 'next/router';
import Link from "next/link";
import { useContext, useState } from "react";
import { AuthContext } from "./_app";

export default function Login() {
  const Auth = useContext(AuthContext)
  const [response, setResponse] = useState();
  const formik = useFormik({
    initialValues: {
      username: "",
      password: "",
    },
    onSubmit: ({ username, password }) => {
      axios.post("http://localhost:3001/login", {username, password}).then(response => { 
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
        <h2>Login</h2>
        <form onSubmit={formik.handleSubmit} className={styles.form}>
          <input id="username" name="username" placeholder="username" type="text" onChange={formik.handleChange} value={formik.values.username} />
          <input id="password" name="password" placeholder="password" type="password" onChange={formik.handleChange} value={formik.values.password} />
          <p className={styles.response}>{response}</p>
          <button type="submit" className={styles.button}>Login</button>
        </form>
        <p className={styles.register}>
          Don&apos;t have an account? <Link href="/register">Register</Link>
        </p>
      </div>
    </div>
  );
}
