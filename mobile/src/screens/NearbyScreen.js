import React, { useEffect, useState } from 'react';
import { View, Text, FlatList, ActivityIndicator, StyleSheet, Image, PermissionsAndroid, Platform, Alert } from 'react-native';
import * as Location from 'expo-location';
import { aiService } from '../services/api';

export default function NearbyScreen({ navigation }) {
  const [location, setLocation] = useState(null);
  const [errorMsg, setErrorMsg] = useState(null);
  const [results, setResults] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    (async () => {
      let { status } = await Location.requestForegroundPermissionsAsync();
      if (status !== 'granted') {
        setErrorMsg('Permission to access location was denied');
        return;
      }

      let location = await Location.getCurrentPositionAsync({});
      setLocation(location);
      fetchNearbyCelebrations(location);
    })();
  }, []);

  const fetchNearbyCelebrations = async (location) => {
    try {
      setLoading(true);
      // We'll create a query that includes the location for nearby search
      // For now, we'll use a simple query near the coordinates
      // In a real app, you might have a dedicated endpoint or use geosearch
      const query = `near:${location.latitude},${location.longitude}`; // Example format
      const response = await aiService.searchCelebrations(query, 10);
      setResults(response.data || []); // Adjust based on actual response structure
    } catch (err) {
      setErrorMsg(`Failed to fetch nearby: ${err.message}`);
      console.error('Nearby search failed:', err);
    } finally {
      setLoading(false);
    }
  };

  const renderItem = ({ item }) => (
    <View style={styles.itemContainer}>
      <Image
        source={{ uri: item.image || 'https://via.placeholder.com/150' }}
        style={styles.thumb}
      />
      <View style={styles.itemDetails}>
        <Text style={styles.itemTitle}>{item.title || 'Untitled'}</Text>
        <Text style={styles.itemLocation}>
          {item.distance?.toFixed(1)} km away {/* Example distance */}
        </Text>
      </View>
    </View>
  );

  if (loading && !location) {
    return (
      <View style={styles.centered}>
        <ActivityIndicator size="large" color="#0000ff" />
      </View>
    );
  }

  if (errorMsg) {
    return (
      <View style={styles.centered}>
        <Text>{errorMsg}</Text>
        <Text onPress={() => {
          // Retry checking permissions
          Location.getForegroundPermissionsAsync().then(({ status }) => {
            if (status === 'granted') {
              Location.getCurrentPositionAsync({}).then(loc => {
                setLocation(loc);
                fetchNearbyCelebrations(loc);
              });
            }
          });
        }}>
          Retry
        </Text>
      </View>
    );
  }

  if (!location) {
    return (
      <View style={styles.centered}>
        <Text>Getting location...</Text>
        <ActivityIndicator size="small" color="#0000ff" />
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Nearby Celebrations</Text>
      {results.length === 0 ? (
        <View style={styles.centered}>
          <Text>No celebrations nearby</Text>
        </View>
      ) : (
        <FlatList
          data={results}
          keyExtractor={item => item.id.toString()}
          renderItem={renderItem}
          contentContainerStyle={styles.listContent}
        />
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 16,
    backgroundColor: '#f5f5f5',
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 20,
    textAlign: 'center',
  },
  centered: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  listContent: {
    paddingBottom: 20,
  },
  itemContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 12,
    marginVertical: 8,
    backgroundColor: '#fff',
    borderRadius: 8,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.2,
    shadowRadius: 1.41,
    elevation: 2,
  },
  thumb: {
    width: 60,
    height: 60,
    borderRadius: 8,
    marginRight: 12,
  },
  itemDetails: {
    flex: 1,
  },
  itemTitle: {
    fontSize: 16,
    fontWeight: '600',
    marginBottom: 4,
  },
  itemLocation: {
    fontSize: 14,
    color: '#666',
  },
});
