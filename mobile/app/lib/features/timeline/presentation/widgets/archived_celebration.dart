import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

class ArchivedCelebration extends StatelessWidget {
  const ArchivedCelebration({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.symmetric(vertical: 10),
      padding: const EdgeInsets.symmetric(horizontal: 10),
      height: 200,
      width: double.maxFinite,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(10),
        color: Colors.white,
      ),
      child: Column(
        children: [
          Row(
            children: [
              Container(
                margin: const EdgeInsets.only(left: 10, top: 10),
                height: 40,
                width: 40,
                decoration: const BoxDecoration(
                  color: Colors.grey,
                  borderRadius: BorderRadius.all(
                    Radius.circular(5),
                  ),
                  image: DecorationImage(
                    fit: BoxFit.cover,
                    image: AssetImage(
                      AppAssets.guyStory,
                    ),
                  ),
                ),
              ),
              const Space(10),
              Text(
                'George Lewis',
                style: context.textTheme.bodyMedium,
              ),
            ],
          ),
          const Space(30),
          Text(
            'Birthday Celebration',
            style: context.textTheme.titleSmall
                ?.copyWith(color: context.colorScheme.primary),
          ),
          Text(
            'We’re interested in your ideas and would be '
            'glad to build something bigger out of it.',
            style: context.textTheme.bodySmall,
            textAlign: TextAlign.center,
          ),
          const Space(20),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Row(
                crossAxisAlignment: CrossAxisAlignment.end,
                children: [
                  Image.asset(
                    AppAssets.handCelebratedIcon,
                    width: 35,
                    height: 35,
                  ),
                  Text(
                    '76',
                    style: context.textTheme.titleSmall?.copyWith(
                      color: context.colorScheme.primary,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                ],
              ),
              Row(
                children: [
                  Image.asset(
                    AppAssets.postComment,
                    width: 22,
                    height: 22,
                    color: context.colorScheme.primary,
                  ),
                  const Space(5),
                  Text(
                    '45',
                    style: context.textTheme.titleSmall?.copyWith(
                      color: context.colorScheme.primary,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                ],
              ),
              Row(
                crossAxisAlignment: CrossAxisAlignment.end,
                children: [
                  SvgPicture.asset(
                    AppAssets.camera,
                    width: 22,
                    height: 22,
                  ),
                  const Space(5),
                  Text(
                    '25',
                    style: context.textTheme.titleSmall?.copyWith(
                      color: context.colorScheme.primary,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                ],
              ),
            ],
          ),
        ],
      ),
    );
  }
}
