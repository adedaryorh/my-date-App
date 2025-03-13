import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class NewPost extends StatefulWidget {
  const NewPost({super.key});

  @override
  State<NewPost> createState() => _NewPostState();
}

class _NewPostState extends State<NewPost> {
  final TextEditingController postCtr = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: false,
      body: SafeArea(
        child: ListView(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                IconButton(
                  onPressed: () {
                    context.pop();
                  },
                  icon: const Icon(
                    Icons.arrow_back_ios,
                    size: 15,
                    color: Colors.black,
                  ),
                ),
                Container(
                  alignment: Alignment.center,
                  height: 30,
                  width: 74,
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10),
                    color: context.colorScheme.primary,
                  ),
                  child: Text(
                    'Speak',
                    style: context.textTheme.bodyMedium
                        ?.copyWith(fontWeight: FontWeight.w600),
                  ),
                ),
              ],
            ),
            const SizedBox(
              height: 30,
            ),
            Row(
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                Container(
                  margin: const EdgeInsets.only(left: 10, top: 10),
                  height: 48,
                  width: 48,
                  decoration: const BoxDecoration(
                    color: Colors.grey,
                    borderRadius: BorderRadius.all(
                      Radius.circular(5),
                    ),
                    image: DecorationImage(
                      fit: BoxFit.cover,
                      image: AssetImage(
                        AppAssets.defaultAvatarGirl,
                      ),
                    ),
                  ),
                ),
                const SizedBox(
                  width: 5,
                ),
                Expanded(
                  child: TextFormInput(
                    controller: postCtr,
                    decoration: InputDecoration(
                      hintText: 'Say something to loved ones',
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
              ],
            ),
          ],
        ),
      ),
    );
  }
}
