import React from "react";
import { createStackNavigator } from "@react-navigation/stack";
import ListingsScreen from "../screens/ListingsScreen";
import ListingDetailsScreen from "../screens/ListingDetailsScreen";
import MyAccountScreen from "../screens/MyAccountScreen";
import MessagesScreen from "../screens/MessagesScreen";
import routes from "./routes";

const Stack = createStackNavigator();

export default function AccountNavigator() {
  return (
    <Stack.Navigator>
      <Stack.Screen name={routes.ACCOUNT} component={MyAccountScreen} />
      <Stack.Screen
        name={routes.MESSAGES}
        component={MessagesScreen}
        options={{ headerShown: false }}
      />
    </Stack.Navigator>
  );
}
