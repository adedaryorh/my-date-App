import 'package:celebut/apps/shared/settings/presentation/widgets/friends_list_item.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class OthersCelebrationPost extends StatefulWidget {
  const OthersCelebrationPost({super.key});

  @override
  State<OthersCelebrationPost> createState() => _OthersCelebrationPostState();
}

class _OthersCelebrationPostState extends State<OthersCelebrationPost> {
  final TextEditingController searchCtr = TextEditingController();
  bool select = false;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          'Celebration Post by Others for me',
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
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          'All Friends',
                          style: context.textTheme.bodyLarge
                              ?.copyWith(fontWeight: FontWeight.w700),
                        ),
                        Transform.scale(
                          scale: 0.7,
                          child: Checkbox(
                            materialTapTargetSize:
                                MaterialTapTargetSize.shrinkWrap,
                            value: select,
                            onChanged: (val) {
                              if (val == null) return;
                              setState(
                                () {
                                  select = val;
                                },
                              );
                            },
                          ),
                        ),
                      ],
                    ),
                    const SizedBox(
                      height: 20,
                    ),
                    ...List.generate(
                      10,
                      (index) => const FriendsListItem(),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
      bottomNavigationBar: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: MainButton(
              loading: false,
              text: 'Save',
              pressed: () {},
            ),
          ),
          const Space(30),
        ],
      ),
    );
  }
}
