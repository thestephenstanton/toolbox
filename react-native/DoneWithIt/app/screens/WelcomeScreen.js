import React from "react";
import { ImageBackground, View, Image, StyleSheet, Text } from "react-native";
import { useNavigation } from "@react-navigation/native";

import Button from "../components/Button";
import colors from "../config/colors";

export default function WelcomeScreen() {
  const navigation = useNavigation();
  return (
    <ImageBackground
      style={styles.background}
      blurRadius={10}
      source={require("../assets/background.jpg")}
    >
      <View style={styles.logoContainer}>
        <Image style={styles.logo} source={require("../assets/logo-red.png")} />
        <Text style={styles.tagline}>Sell What You Don't Need</Text>
      </View>

      <View style={styles.buttonsContainer}>
        <Button
          title="login"
          color={colors.primary}
          onPress={() => navigation.navigate("Login")}
        />
        <Button
          title="register"
          color={colors.secondary}
          onPress={() => navigation.navigate("Register")}
        />
      </View>
    </ImageBackground>
  );
}

const styles = StyleSheet.create({
  background: {
    flex: 1,
    justifyContent: "flex-end",
    alignItems: "center",
  },
  logo: {
    height: 100,
    width: 100,
  },
  logoContainer: {
    position: "absolute",
    top: 70,
    alignItems: "center",
  },
  tagline: {
    fontSize: 25,
    fontWeight: "600",
    paddingVertical: 20,
  },
  buttonsContainer: {
    padding: 20,
    width: "100%",
  },
});
