import styles from "./page.module.css";
import { ApiButton } from "./components/apiButton";

export default function Home() {

  return (
    <div className={styles.page}>
      <main className={styles.main}>
        <ApiButton />
      </main>
      <footer className={styles.footer}>

      </footer>
    </div>
  );
}
