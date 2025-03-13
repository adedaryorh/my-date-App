import 'package:freezed_annotation/freezed_annotation.dart';

part 'generated/post_model.freezed.dart';
part 'generated/post_model.g.dart';

@freezed
class PostModel with _$PostModel {
  const factory PostModel({
    @Default('') String? userName,
    @Default('') String? time,
    @Default('') String? content,
    @Default([]) List<String> media,
    @Default(false) bool isVideo,
  }) = _PostModel;

  factory PostModel.fromJson(Map<String, Object?> json) =>
      _$PostModelFromJson(json);
}
