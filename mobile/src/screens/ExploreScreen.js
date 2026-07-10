import React, { useState } from 'react';
import { View, Text, TextInput, FlatList, ActivityIndicator, StyleSheet, Image, TouchableOpacity } from 'react-native';
import { aiService } from '../services/api';

export default function ExploreScreen({ navigation }) {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const searchCelebrations = async () => {
    if (!query.trim()) {
      setResults([]);
      return;
    }

    try {
      setLoading(true);
      const response = await aiService.searchCelebrations(query, 10);
      setResults(response.data || []); // Adjust based on actual response structure
    } catch (err) {
      setError(err.message);
      console.error('Search failed:', err);
    } finally {
      setLoading(false);
    }
  };

  const renderItem = ({ item }) => (
    <View style={styles.itemContainer}>
      {/* Assuming item has properties like id, title, image, etc. */}
      <Image
        source={{ uri: item.image || 'https://via.placeholder.com/150' }}
        style={styles.thumb}
      />
      <View style={styles.itemDetails}>
        <Text style={styles.itemTitle}>{item.title || 'Untitled'}</Text>
        <Text style={styles.itemDescription}>{item.description || ''}</Text>
      </View>
    </View>
  );

  if (loading) {
    return (
      <View style={styles.centered}>
        <ActivityIndicator size="large" color="#0000ff" />
      </View>
    );
  }

  if (error) {
    return (
      <View style={styles.centered}>
        <Text>Error: {error}</Text>
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Explore Celebrations</Text>
      <View style={styles.searchContainer}>
        <TextInput
          style={styles.searchInput}
          placeholder="Search celebrations..."
          value={query}
          onChangeText={setQuery}
          onSubmitEditing={searchCelebrations}
        />
        <TouchableOpacity style={styles.searchButton} onPress={searchCelebrations}>
          <Text style={styles.searchButtonText}>Search</Text>
        </TouchableOpacity>
      </View>

      {results.length === 0 && !loading && !query ? (
        <View style={styles.centered}>
          <Text>Enter a search term to find celebrations</Text>
        </View>
      ) : results.length === 0 && !loading ? (
        <View style={styles.centered}>
          <Text>No results found</Text>
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
  searchContainer: {
    flexDirection: 'row',
    marginBottom: 20,
  },
  searchInput: {
    flex: 1,
    height: 40,
    borderColor: '#ccc',
    borderWidth: 1,
    borderRadius: 4,
    paddingHorizontal: 12,
    marginRight: 8,
  },
  searchButton: {
    backgroundColor: '#007aff',
    paddingVertical: 8,
    paddingHorizontal: 16,
    borderRadius: 4,
  },
  searchButtonText: {
    color: '#fff',
    fontWeight: '600',
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
  itemDescription: {
    fontSize: 14,
    color: '#666',
  },
});
