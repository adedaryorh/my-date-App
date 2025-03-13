import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class FriendItem extends StatelessWidget {
  const FriendItem({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 56,
      width: double.maxFinite,
      margin: const EdgeInsets.only(bottom: 10),
      decoration: const BoxDecoration(
        color: Color(0xffF4F4F4),
        borderRadius: BorderRadius.all(Radius.circular(10)),
      ),
      child: Row(
        children: [
          Container(
            height: 48,
            width: 48,
            padding: const EdgeInsets.all(2),
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(16),
              color: Colors.grey.shade300,
              image: const DecorationImage(
                fit: BoxFit.cover,
                image: AssetImage(AppAssets.girlStory),
              ),
            ),
          ),
          const Space(10),
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text('Tobi Ozenua'),
              Text(
                '@tobiozenua',
                style: context.textTheme.bodySmall?.copyWith(
                  color: const Color(0xffADB5BD),
                ),
              ),
            ],
          )
        ],
      ),
    );
  }
}
