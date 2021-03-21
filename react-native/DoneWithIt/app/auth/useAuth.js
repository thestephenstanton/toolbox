import jwtDecode from "jwt-decode";
import { useContext } from "react";

import AuthContext from "./context";
import authStorage from "./storage";

const useAuth = () => {
  const { user, setUser } = useContext(AuthContext);

  const logIn = (jwt) => {
    const user = jwtDecode(jwt);
    setUser(user);
    authStorage.storeJwt(jwt);
  };

  const logOut = () => {
    setUser(null);
    authStorage.removeJwt();
  };

  return { user, setUser, logIn, logOut };
};

export default useAuth;
