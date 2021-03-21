import React from "react";
import LottieView from "lottie-react-native";
import { View, StyleSheet } from "react-native";

export default function ActivityIndicator({ isVisible }) {
  if (!isVisible) return null;
  return (
    <View style={styles.container}>
      <LottieView
        autoPlay
        loop
        source={require("../assets/animations/loader.json")}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    height: "100%",
    width: "100%",
    backgroundColor: "white",
    position: "absolute",
    zIndex: 1,
    opacity: 0.8,
  },
});
