import 'package:celebut/apps/shared/settings/presentation/widgets/privacy_option_item.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class PrivacyAndSafety extends StatefulWidget {
  const PrivacyAndSafety({super.key});

  @override
  State<PrivacyAndSafety> createState() => _PrivacyAndSafetyState();
}

class _PrivacyAndSafetyState extends State<PrivacyAndSafety> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const SizedBox(
                height: 50,
              ),
              InkWell(
                onTap: () {
                  context.pop();
                },
                child: const SizedBox(
                  height: 20,
                  width: 20,
                  child: Icon(
                    Icons.arrow_back_ios,
                    size: 13,
                  ),
                ),
              ),
              const SizedBox(
                height: 20,
              ),
              Text(
                'Privacy and Safety',
                style: context.textTheme.headlineMedium,
              ),
              const Text('Chose who & what can seen on your timeline'),
              const Space(50),
              ...List.generate(
                AppConstants.privacyOptions.length,
                (index) => PrivacyOptionItem(
                  title: AppConstants.privacyOptions[index],
                  tapped: () {},
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
