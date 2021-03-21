import { useNavigation } from "@react-navigation/native";
import React, { useContext } from "react";
import { View, StyleSheet } from "react-native";
import useAuth from "../auth/useAuth";
import ListItem from "../components/ListItem";
import ListItemSeparator from "../components/ListItemSeparator";
import Screen from "../components/Screen";
import ShortcutButton from "../components/ShortcutButton";
import colors from "../config/colors";
import routes from "../navigation/routes";

export default function MyAccountScreen() {
  const navigation = useNavigation();
  const { user, logOut } = useAuth();

  return (
    <Screen style={styles.container}>
      <View style={styles.userContainer}>
        <ListItem
          title={user.name}
          subTitle={user.email}
          image={require("../assets/mosh.jpg")}
        />
      </View>
      <View style={styles.options1Container}>
        <ShortcutButton
          text="My Listings"
          icon="format-list-bulleted"
          color={colors.primary}
          onPress={() => {}}
        />
        <ListItemSeparator />
        <ShortcutButton
          text="My Messages"
          icon="email"
          color={colors.secondary}
          onPress={() => navigation.navigate(routes.MESSAGES)}
        />
      </View>
      <View style={styles.options2Container}>
        <ShortcutButton
          text="Log Out"
          icon="logout"
          color={colors.yellow}
          onPress={logOut}
        />
      </View>
    </Screen>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.light,
  },
  userContainer: {
    backgroundColor: colors.white,
  },
  options1Container: {
    paddingTop: 40,
  },
  options2Container: {
    paddingTop: 20,
  },
});
