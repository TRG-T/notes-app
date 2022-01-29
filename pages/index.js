import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";

export default function Home() {
  return (
    <div className={styles.container}>
      <div className={styles.login}>
        <h2>Login</h2>
        <div className={styles.inputs}>
          <input placeholder="nickname" type="text" />
          <input placeholder="password" type="password"/>
        </div>
        <button className={styles.button}>Login</button>
        <p className={styles.register}>Don&apos;t have an account? <a>Register</a></p>
      </div>
    </div>
  );
}
