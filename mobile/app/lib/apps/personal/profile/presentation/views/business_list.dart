import 'package:celebut/apps/personal/profile/presentation/widgets/friend_item.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class BusinessList extends StatefulWidget {
  const BusinessList({super.key});

  @override
  State<BusinessList> createState() => _BusinessListState();
}

class _BusinessListState extends State<BusinessList> {
  final TextEditingController searchCtr = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        centerTitle: true,
        title: Text(
          'Your Businesses',
          style: context.textTheme.titleSmall,
        ),
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
              const Space(20),
              AppSearchBar(
                searchCtr: searchCtr,
                labelText: 'Search',
              ),
              const Space(20),
              Expanded(
                child: ListView(
                  shrinkWrap: true,
                  children: [
                    ...List.generate(
                      10,
                      (index) => const FriendItem(),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
