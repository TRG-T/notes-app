import '../styles/globals.css'

import { createContext, useEffect, useState } from "react"
const AuthContext = createContext()

function MyApp({ Component, pageProps }) {
  const [auth, setAuth] = useState();
  useEffect(() => {
    setAuth(false)
  }, [])
  return ( 
    <AuthContext.Provider value={{auth, setAuth}}>
      <Component {...pageProps} />
    </AuthContext.Provider>
  )
}

export default MyApp
export { AuthContext }