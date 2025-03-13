import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class BlockedAccounts extends StatefulWidget {
  const BlockedAccounts({super.key});

  @override
  State<BlockedAccounts> createState() => _BlockedAccountsState();
}

class _BlockedAccountsState extends State<BlockedAccounts> {
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
                'Blocked Accounts',
                style: context.textTheme.headlineMedium,
              ),
              const Spacer(),
              Align(
                child: Column(
                  children: [
                    Text(
                      'No Blocked Accounts',
                      style: context.textTheme.titleMedium,
                    ),
                    const Space(20),
                    Text(
                      'You have no blocked accounts',
                      style: context.textTheme.bodyMedium,
                    ),
                  ],
                ),
              ),
              const Spacer(),
            ],
          ),
        ),
      ),
    );
  }
}
