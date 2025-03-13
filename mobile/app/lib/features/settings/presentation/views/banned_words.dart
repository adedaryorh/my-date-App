import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class BannedWords extends StatefulWidget {
  const BannedWords({super.key});

  @override
  State<BannedWords> createState() => _BannedWordsState();
}

class _BannedWordsState extends State<BannedWords> {
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
                'Banned Words',
                style: context.textTheme.headlineMedium,
              ),
              const Text('Ban words you don’t want to see on your timeline'),
              const Spacer(),
              Column(
                children: [
                  Text(
                    'No Banned Words',
                    style: context.textTheme.titleMedium,
                  ),
                  const Text(
                    'You’ve not banned any words. '
                    'Select words you will like to ban now',
                    textAlign: TextAlign.center,
                  ),
                  TextButton(
                    onPressed: () {
                      context.showBannedWordsDialog();
                    },
                    child: Text(
                      'Add Banned Words',
                      style: context.textTheme.bodySmall
                          ?.copyWith(color: context.colorScheme.primary),
                    ),
                  ),
                ],
              ),
              const Spacer(),
            ],
          ),
        ),
      ),
    );
  }
}
