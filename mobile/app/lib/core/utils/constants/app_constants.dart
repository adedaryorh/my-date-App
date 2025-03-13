import 'package:celebut/core/core.dart';

class AppConstants {
  static const onboardColors = [
    AppTheme.onboard1,
    AppTheme.onboard2,
    AppTheme.onboard3,
  ];

  static const industryTypes = [
    'Entertainment',
    'Agriculture',
    'Finance',
    'Mining',
    'Information Technology',
    'Furniture',
  ];

  static const settingOptions = [
    'Interests',
    'Notifications',
    'Language',
    'Change Password',
    'Request for Verification',
    'Linked Devices',
    'Celebration Post by Others for me',
    'Surprise Event by Others for me',
    'Privacy and safety',
    'Log out',
  ];

  static const interestOptions = [
    'News and Events',
    'Entertainment',
    'Lifestyle',
    'Personal Development',
    'Humor and Memes',
    'Sports',
    'Science',
    'History',
    'Animals',
    'Education',
    'Technology',
    'Product and Brand',
    'Marketing',
    'Scary Things',
    'Movies',
    'Movies',
  ];

  static const Map<String, String> languageOptions = {
    '🇮🇩': 'Indonesia',
    '🇺🇸': 'English',
    '🇹🇭': 'Thailand',
    '🇨🇳': 'Chinese',
  };

  static const privacyOptions = ['Content', 'Banned Words', 'Blocked Accounts'];

  static const notificationOptions = [
    'Likes',
    'Comments',
    'Tags & Mentions',
    'Repost',
    'Direct Message',
    'Live',
    'New Followers ',
  ];

  static const reportOptions = [
    'Child Sexual Abuse Material',
    'Extremism',
    'Drug',
    'Gambling',
    'Pornography',
    'Graphic Violence',
    'Weapon',
  ];

  static const Map<String, String> profileDetailOptions = {
    AppAssets.profilePerson: 'Personal Data',
    AppAssets.profileWallet: 'Wallet',
    AppAssets.settings2: 'Notifications',
  };

  static const Map<String, String> profileSupportOptions = {
    AppAssets.profileInfo: 'Help Center',
    AppAssets.profileAddPerson: 'Add another account',
  };

  static const featureList = [
    'Wishlist',
    'Access to premium celebrations feature',
    'Extended celebration posts',
    'Revenue from celebrations',
    'Write longer posts and articles when sharing updates on your timeline.',
  ];

  static const whoCanSeeContent = [
    'Everyone',
    'Friends Only',
    'Friends & Friends of Friends',
    'Selected Friends',
  ];
}
