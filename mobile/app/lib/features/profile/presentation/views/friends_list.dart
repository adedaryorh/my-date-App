import 'package:celebut/core/core.dart';
import 'package:celebut/features/profile/presentation/widgets/friend_item.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class FriendsList extends StatefulWidget {
  const FriendsList({super.key});

  @override
  State<FriendsList> createState() => _FriendsListState();
}

class _FriendsListState extends State<FriendsList> {
  final TextEditingController searchCtr = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        centerTitle: true,
        title: Text(
          'Your friends',
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
