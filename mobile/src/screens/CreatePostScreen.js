import React, { useState } from 'react';
import { View, Text, TextInput, Button, ActivityIndicator, StyleSheet, Alert, Image } from 'react-native';
import { aiService } from '../services/api';

export default function CreatePostScreen({ navigation }) {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [imageUrl, setImageUrl] = useState(''); // In a real app, you'd have image upload
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async () => {
    if (!title.trim() && !description.trim()) {
      setError('Please enter a title or description');
      return;
    }

    setLoading(true);
    setError('');

    try {
      // First, moderate the content (title + description)
      const contentToCheck = `${title} ${description}`.trim();
      if (contentToCheck) {
        const moderationResult = await aiService.moderateCelebration(contentToCheck);
        
        // Assuming the moderation result has a flag indicating if it's safe
        // Adjust based on actual response from your AI service
        if (moderationResult.data && moderationResult.data.isToxic) { // Example field
          throw new Error('Content contains inappropriate language and cannot be posted.');
        }
      }

      // If we get here, content is approved. Now create the post.
      // This is a placeholder for the actual API call to create a celebration
      // You would replace this with a call to your backend's celebration creation endpoint
      console.log('Creating post with:', { title, description, imageUrl });
      
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 1500));
      
      Alert.postedSuccessfully());
      // Reset form
      setTitle('');
      setDescription('');
      setImageUrl('');
      navigation.goBack(); // Go back to previous screen
    } catch (err) {
      setError(err.message || 'Failed to create post');
      console.error('Create post error:', err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Create a Celebration</Text>
      
      {error && (
        <View style={styles.error}>
          <Text>{error}</Text>
        </View>
      )}
      
      <TextInput
        style={styles.input}
        placeholder="Title"
        value={title}
        onChangeText={setTitle}
        autoCapitalize="words"
      />
      
      <TextInput
        style={styles.input}
        placeholder="Description"
        value={description}
        onChangeText={setDescription}
        autoCapitalize="sentences"
        multiline
        minHeight={80}
      />
      
      {/* In a real app, you'd have an image picker here */}
      <View style={styles.inputContainer}>
        <Text style={styles.label}>Image URL (optional):</Text>
        <TextInput
          style={styles.input}
          placeholder="https://example.com/image.jpg"
          value={imageUrl}
          onChangeText={setImageUrl}
        />
      </View>
      
      <Button
        title="Post Celebration"
        onPress={handleSubmit}
        disabled={loading}
        color="#007aff"
      />
      
      {loading && (
        <View style={styles.loadingContainer}>
          <ActivityIndicator size="large" color="#007aff" />
          <Text style={styles.loadingText}>Posting...</Text>
        </View>
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
    marginBottom: 24,
    textAlign: 'center',
  },
  input: {
    height: 40,
    borderColor: '#ccc',
    borderWidth: 1,
    borderRadius: 4,
    paddingHorizontal: 12,
    marginBottom: 12,
    backgroundColor: '#fff',
  },
  inputContainer: {
    marginBottom: 12,
  },
  label: {
    fontSize: 14,
    marginBottom: 4,
    color: '#555',
  },
  error: {
    backgroundColor: '#ffebee',
    borderColor: '#f44336',
    borderWidth: 1,
    borderRadius: 4,
    padding: 12,
    marginBottom: 16,
  },
  errorText: {
    color: '#c62828',
  },
  loadingContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    paddingVertical: 20,
  },
  loadingText: {
    marginLeft: 12,
    fontSize: 16,
    color: '#666',
  },
});
