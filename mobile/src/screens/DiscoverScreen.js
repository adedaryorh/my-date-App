import React, { useEffect, useState } from 'react';
import { View, Text, FlatList, ActivityIndicator, StyleSheet, Image, TouchableOpacity } from 'react-native';
import { aiService } from '../services/api';

export default function DiscoverScreen({ navigation }) {
  const [recommendations, setRecommendations] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    loadRecommendations();
  }, []);

  const loadRecommendations = async () => {
    try {
      setLoading(true);
      // In a real app, you would get the userId from auth context or storage
      const response = await aiService.getUserRecommendations(10);
      setRecommendations(response.recommended_users || []);
    } catch (err) {
      setError(err.message);
      console.error('Failed to load recommendations:', err);
    } finally {
      setLoading(false);
    }
  };

  const renderItem = ({ item }) => (
    <View style={styles.itemContainer}>
      {/* Assuming item has properties like id, name, avatar, etc. */}
      <Image
        source={{ uri: item.profile_image_url || 'https://via.placeholder.com/150' }}
        style={styles.avatar}
      />
      <View style={styles.itemDetails}>
        <Text style={styles.itemName}>{`${item.first_name || ''} ${item.last_name || ''}`.trim() || item.username}</Text>
        <Text style={styles.itemBio}>{item.bio || ''}</Text>
      </View>
      <TouchableOpacity style={styles.followButton} onPress={() => handleFollow(item.user_id)}>
        <Text>{item.isFollowing ? 'Following' : 'Follow'}</Text>
      </TouchableOpacity>
    </View>
  );

  const handleFollow = async (userIdToFollow) => {
    try {
      await aiService.logInteraction(
        userIdToFollow,
        'follow'
      );
      // Update the item's state optimistically
      setRecommendations(prev =>
        prev.map(item =>
          item.user_id === userIdToFollow ? { ...item, isFollowing: true } : item
        )
      );
    } catch (err) {
      console.error('Failed to follow user:', err);
    }
  };

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
      <Text style={styles.title}>Discover People</Text>
      {recommendations.length === 0 ? (
        <View style={styles.centered}>
          <Text>No recommendations available</Text>
        </View>
      ) : (
        <FlatList
          data={recommendations}
          keyExtractor={item => item.user_id}
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
  avatar: {
    width: 50,
    height: 50,
    borderRadius: 25,
    marginRight: 12,
  },
  itemDetails: {
    flex: 1,
  },
  itemName: {
    fontSize: 16,
    fontWeight: '600',
    marginBottom: 4,
  },
  itemBio: {
    fontSize: 14,
    color: '#666',
  },
  followButton: {
    backgroundColor: '#007aff',
    paddingVertical: 8,
    paddingHorizontal: 16,
    borderRadius: 4,
  },
});
