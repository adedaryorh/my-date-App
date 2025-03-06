import 'package:celebut/apps/shared/settings/presentation/widgets/friends_list_item.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class SharePost extends StatefulWidget {
  const SharePost({super.key});

  @override
  State<SharePost> createState() => _SharePostState();
}

class _SharePostState extends State<SharePost> {
  final TextEditingController searchCtr = TextEditingController();
  final TextEditingController messageCtr = TextEditingController();
  bool select = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          'Share post',
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
          const Space(10),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: TextFormInput(
              controller: messageCtr,
              decoration: InputDecoration(
                hintText: 'Write a message',
                hintStyle: context.textTheme.bodyMedium?.copyWith(
                  fontWeight: FontWeight.w400,
                  color: const Color(0xffBDBDBD),
                ),
                isDense: true,
                border: OutlineInputBorder(
                  borderSide: const BorderSide(color: Color(0xffF3F3F3)),
                  borderRadius: BorderRadius.circular(7),
                ),
                focusedBorder: OutlineInputBorder(
                  borderSide: const BorderSide(color: Color(0xffF3F3F3)),
                  borderRadius: BorderRadius.circular(7),
                ),
                enabledBorder: OutlineInputBorder(
                  borderSide: const BorderSide(color: Color(0xffF3F3F3)),
                  borderRadius: BorderRadius.circular(7),
                ),
              ),
            ),
          ),
          const Space(10),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: MainButton(
              loading: false,
              text: 'Send',
              pressed: () {},
            ),
          ),
          const Space(30),
        ],
      ),
    );
  }
}
