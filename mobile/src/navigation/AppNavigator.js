import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';

import DiscoverScreen from '../screens/DiscoverScreen';
import ExploreScreen from '../screens/ExploreScreen';
import NearbyScreen from '../screens/NearbyScreen';
import CreatePostScreen from '../screens/CreatePostScreen';

const Stack = createStackNavigator();

export default function AppNavigator() {
  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName="Discover">
        <Stack.Screen name="Discover" component={DiscoverScreen} />
        <Stack.Screen name="Explore" component={ExploreScreen} />
        <Stack.Screen name="Nearby" component={NearbyScreen} />
        <Stack.Screen name="CreatePost" component={CreatePostScreen} />
        {/* Add more screens as needed */}
      </Stack.Navigator>
    </NavigationContainer>
  );
}
