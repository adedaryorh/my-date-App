import 'package:celebut/core/core.dart';
import 'package:celebut/features/profile/presentation/widgets/profile_avatar.dart';
import 'package:celebut/features/profile/presentation/widgets/profile_detail_item.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ProfileDetails extends StatefulWidget {
  const ProfileDetails({super.key});

  @override
  State<ProfileDetails> createState() => _ProfileDetailsState();
}

class _ProfileDetailsState extends State<ProfileDetails> {
  @override
  Widget build(BuildContext context) {
    const profileDetailOptions = AppConstants.profileDetailOptions;
    const supportOptions = AppConstants.profileSupportOptions;
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          onPressed: () {
            context.pop();
          },
          icon: const Icon(
            Icons.arrow_back_ios,
            size: 15,
          ),
        ),
      ),
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: Column(
            children: [
              const Space(80),
              const ProfileAvatar(),
              const Space(20),
              Text(
                'Alice Smith',
                style: context.textTheme.titleSmall
                    ?.copyWith(fontWeight: FontWeight.w700),
              ),
              const Text('Alicesmith@gmail.com'),
              const Space(20),
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'Profile',
                  style: context.textTheme.bodySmall,
                ),
              ),
              const Space(20),
              ...List.generate(
                profileDetailOptions.keys.toList().length,
                (index) {
                  final assetPaths = profileDetailOptions.keys.toList();
                  final titles = profileDetailOptions.values.toList();
                  return ProfileDetailItem(
                    assetPath: assetPaths[index],
                    title: titles[index],
                    tapped: () {
                      switch (titles[index]) {
                        case 'Personal Data':
                          context.pushNamed(AppRoute.pEditProfile.name);
                        case 'Wallet':
                          break;
                        case 'Notifications':
                          break;
                      }
                    },
                  );
                },
              ),
              const Space(20),
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'Support',
                  style: context.textTheme.bodySmall,
                ),
              ),
              const Space(20),
              ...List.generate(
                supportOptions.keys.toList().length,
                (index) {
                  final assetPaths = supportOptions.keys.toList();
                  final titles = supportOptions.values.toList();
                  return ProfileDetailItem(
                    assetPath: assetPaths[index],
                    title: titles[index],
                    tapped: () {},
                  );
                },
              ),
            ],
          ),
        ),
      ),
    );
  }
}
