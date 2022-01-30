import styles from "../styles/Home.module.css";
import axios from "axios"
import { useFormik } from "formik";
import Router from 'next/router';
import { useContext } from "react";
import { AuthContext } from "./_app";

export default function Home() {
  const Auth = useContext(AuthContext)
  const formik = useFormik({
    initialValues: {
      username: "",
      password: "",
    },
    onSubmit: ({ username, password }) => {
      axios.post("http://localhost:3001/login", {username, password}).then(response => { 
        console.log(response.data)
        if (response.data === "Success") {
          Router.push("/home")
          Auth.setAuth(true)
        }
      })
    }
  });

  return (
    <div className={styles.container}>
      <div className={styles.login}>
        <h2>Login</h2>
        <form onSubmit={formik.handleSubmit} className={styles.inputs}>
          <input id="username" name="username" placeholder="username" type="text" onChange={formik.handleChange} value={formik.values.username} />
          <input id="password" name="password" placeholder="password" type="password" onChange={formik.handleChange} value={formik.values.password} />
          <button type="submit" className={styles.button}>Login</button>
        </form>
        <p className={styles.register}>
          Don&apos;t have an account? <a>Register</a>
        </p>
      </div>
    </div>
  );
}
